<template>
 <ul id="v-for-object" class="demo">
  <li v-for="value in jobs">
    {{ value }}
  </li>
</ul>
</template>

<script>
import axios from "axios";
export default {
  data() {
    return {
      foo: "hoge",
      jobs: []
    };
  },
  mounted() {
    axios.get("https://hal-iot.net/api/admin/jobs").then(res => {
      this.jobs = res.data.jobs;
      this.jobs.reverse();
      this.jobs = this.jobs.slice(0, 5);
    });
  }
};
</script>

<style>
</style>

