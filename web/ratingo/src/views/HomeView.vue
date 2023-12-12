<script setup lang="ts">
import router from '../router/index';
import { ref } from 'vue';
import { config } from '../config/config';
import axios from 'axios';

interface Rating {
  topic: string;
  name: string;
  mail: string;
  rating: number;
  comment: string;
}

const topic = ref('');
const name = ref('');
const mail = ref('');
const rate = ref(0);
const comment = ref('');
const fullscreenLoading = ref(false);

const submit =async () => {
  fullscreenLoading.value = true;
  const req: Rating = {
    topic: topic.value,
    name: name.value,
    mail: mail.value,
    rating: rate.value,
    comment: comment.value,
  };
  const res = await axios.post(`api/collect`, req);
  console.log(res);
  if (res.status === 200) {
    // res body: {Success:true ID:1 TopicID:1}
    const res2 = await axios.post(`api/analyze`, { topic: res.data.topic_id });
    if (res2.status === 200) {
      fullscreenLoading.value = false;
      router.push('/rating?topic=' + res.data.topic_id);
    } else {
      fullscreenLoading.value = false;
      alert('Submit failed!');
    }
  } else {
    fullscreenLoading.value = false;
    alert('Submit failed!');
  }
}

</script>

<template>
  <main>
    <h1>Ratingo</h1>
    <el-card>
      <el-form label-width="80px">
        <el-form-item label="Topic" required>
          <el-input v-model="topic" placeholder="Topic"></el-input>
        </el-form-item>
        <el-form-item label="Name" required>
          <el-input v-model="name" placeholder="Your Name"></el-input>
        </el-form-item>
        <el-form-item label="Mail" required>
          <el-input v-model="mail" placeholder="Your Mail"></el-input>
        </el-form-item>
        <el-form-item label="Rating" required>
          <el-rate v-model="rate" allow-half></el-rate>
        </el-form-item>
        <el-form-item label="Comment" required>
          <el-input type="textarea" v-model="comment" placeholder="Comment"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button v-loading.fullscreen.lock="fullscreenLoading"
          type="primary" @click="submit">Submit</el-button>
        </el-form-item>
      </el-form>
    </el-card>
    <el-card>
      <p>
        Ratingo is for the CSCA5028 final project.All code is available on <a
          href="github.com/mizumoto-cn/ratingo">GitHub</a>.
      </p>
      <p>
        This project is governed by the <a href="https://github.com/mizumoto-cn/ratingo/tree/main/LICENCE">MGPL
          License</a>(Mizumoto.General.Public.License.v1.4).
      </p>
      <p>
        Copyright@mizumoto-cn 2023. All rights reserved.
      </p>
    </el-card>
  </main>
</template>
