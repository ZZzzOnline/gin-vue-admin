<template>
    <div class="waterflow">
        <div v-for="(image, index) in images" :key="index" class="waterflow-item">
            <img :src="image.src" :alt="image.alt" />
        </div>
    </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted } from 'vue';
import axios from 'axios';

const images = ref([]);
const imageHashes = ref({});

// 从 LocalStorage 加载缓存的图片元数据和哈希值
const loadCachedImages = () => {
    const cachedImages = localStorage.getItem('cachedImages');
    const cachedHashes = localStorage.getItem('cachedHashes');
    if (cachedImages && cachedHashes) {
        images.value = JSON.parse(cachedImages);
        imageHashes.value = JSON.parse(cachedHashes);
    }
};

// 将图片元数据和哈希值缓存到 LocalStorage
const cacheImages = () => {
    localStorage.setItem('cachedImages', JSON.stringify(images.value));
    localStorage.setItem('cachedHashes', JSON.stringify(imageHashes.value));
};

// 从服务器加载图片元数据
const fetchImages = async () => {
    try {
        const response = await axios.get('http://localhost:8080/api/images', {
            headers: { Authorization: 'valid-token' }
        });
        const newImages = response.data.map(img => ({
            src: `http://localhost:8080${img.url}`,
            alt: img.description,
            hash: img.hash
        }));

        // 检查哈希值是否变化
        let updated = false;
        newImages.forEach(img => {
            if (imageHashes.value[img.src] !== img.hash) {
                updated = true;
                imageHashes.value[img.src] = img.hash;
            }
        });

        if (updated) {
            images.value = newImages;
            cacheImages();
        }
    } catch (error) {
        console.error('Error fetching images:', error);
    }
};

const loadMoreImages = async () => {
    try {
        const response = await axios.get('http://localhost:8080/api/images?page=2', {
            headers: { Authorization: 'valid-token' }
        });
        const newImages = response.data.map(img => ({
            src: `http://localhost:8080${img.url}`,
            alt: img.description,
            hash: img.hash
        }));

        // 检查哈希值是否变化
        let updated = false;
        newImages.forEach(img => {
            if (imageHashes.value[img.src] !== img.hash) {
                updated = true;
                imageHashes.value[img.src] = img.hash;
            }
        });

        if (updated) {
            images.value.push(...newImages);
            cacheImages();
        }
    } catch (error) {
        console.error('Error loading more images:', error);
    }
};

const handleScroll = () => {
    if ((window.innerHeight + window.scrollY) >= document.body.offsetHeight - 100) {
        loadMoreImages();
    }
};

onMounted(() => {
    loadCachedImages();
    if (images.value.length === 0) {
        fetchImages();
    }
    window.addEventListener('scroll', handleScroll);
});

onUnmounted(() => {
    window.removeEventListener('scroll', handleScroll);
});

const eventSource = new EventSource("/events");
eventSource.onmessage = function (event) {
    const newElement = document.createElement("div");
    newElement.innerHTML = "New message: " + event.data;
    document.getElementById("events").appendChild(newElement);
};

</script>

<style scoped>
.waterflow {
    column-count: 3;
    column-gap: 1em;
}

.waterflow-item {
    break-inside: avoid;
    margin-bottom: 1em;
}

.waterflow-item img {
    width: 100%;
    display: block;
}
</style>