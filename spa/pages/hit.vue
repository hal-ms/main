<template>
<img :src="img">
</template>

<script>
export default {
  data() {
    return {
      img:
        "http://static-maquia.hpplus.jp/upload/image/manager/25/EGIXA1Q-1200.jpg"
    };
  },
  mounted() {}
};
</script>

<style>
img {
  width: 100%;
  margin: 0 auto;
  display: block;
}
</style>
