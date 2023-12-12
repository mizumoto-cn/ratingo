<template>
    <div>
        <el-card>
            <h2>Ratingo Topic Detail</h2>
            <p style="margin-top: 20px;margin-right: 0%;text-align: right;">
                <el-button type="primary" @click="back">Back</el-button>
            </p>
            <p>Topic: {{ topic }}</p>
            <p>Average Rating: </p>
            <el-rate v-model="value" disabled show-score text-color="#ff9900" score-template="{value} points" />
            <p>Rating Count: {{ ratingCount }}</p>
            <el-table :data="details" stripe>
                <el-table-column prop="name" label="Name" width="180"></el-table-column>
                <el-table-column prop="rating" label="Rating" width="180"></el-table-column>
                <el-table-column prop="comment" label="Comment"></el-table-column>
            </el-table>
            <el-pagination
                @size-change="handleSizeChange"
                @current-change="handleCurrentChange"
                :current-page="currentPage"
                :page-size="pageSize"
                layout="total, sizes, prev, pager, next, jumper"
                :page-sizes="[10, 20, 30, 40]"
                :total="details.length">
            </el-pagination>
        </el-card>
    </div>
</template>

<script setup lang="ts">
import axios from 'axios';
import { ref, onMounted, computed } from 'vue';

const topic = ref('');
const value = ref(0);
const ratingCount = ref(0.0);
const details = ref<RInfo['details']>([]);

// {"avg":4.318182,"details":[{"name":"1****1","rating":3.5,"comment":"1"},{"name":"1****1","rating":4.5,"comment":"1"},{"name":"1****1","rating":4.5,"comment":"1"},{"name":"1****1","rating":4.5,"comment":"1"},{"name":"1****1","rating":4.5,"comment":"1"},{"name":"1****1","rating":4.5,"comment":"1"},{"name":"1****1","rating":4.5,"comment":"1"},{"name":"1****1","rating":4.5,"comment":"1"},{"name":"1****1","rating":4.5,"comment":"1"},{"name":"1****1","rating":4.5,"comment":"1"},{"name":"1****1","rating":3.5,"comment":"1"}]}
interface RInfo {
    topic: string;
    avg: number;
    details: {
        name: string;
        rating: number;
        comment: string;
    }[];
}

onMounted(async () => {
    const url = new URL(window.location.href);
    const topicId = url.searchParams.get('topic') || null;
    if (topicId) {
        const tid = parseInt(topicId);
        await getTopicInfo(tid);
    }
});

const getTopicInfo = async (id: number) => {
    const res = await axios.post(`api/display`, { topic: id })
    if (res.status === 200) {
        const res2: RInfo = res.data;
        topic.value = res2.topic;
        value.value = res2.avg;
        ratingCount.value = res2.details.length;
        details.value = res2.details;
    } else {
        alert('Get topic info failed!');
    }
}

const currentPage = ref(1);
const pageSize = ref(10);
const pagedDetails = computed(() => {
    const start = (currentPage.value - 1) * pageSize.value;
    const end = start + pageSize.value;
    return details.value.slice(start, end);
});

const handleSizeChange = (val: number) => {
    pageSize.value = val;
    currentPage.value = 1; // Reset to first page when page size changes
};

const handleCurrentChange = (val: number) => {
    currentPage.value = val;
};

</script>

<style scoped></style>