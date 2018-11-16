<template>
<img :src="img">
</template>

<script>
import io from "socket.io-client";
export default {
  data() {
    return {
      img:
        "http://static-maquia.hpplus.jp/upload/image/manager/25/EGIXA1Q-1200.jpg"
    };
  },
  mounted() {
    this.socket = io("https://socket.patra.store", {
      transports: ["websocket"]
    });
    this.socket.on("hit", msg => {
      this.img = JSON.parse(msg).img;
    });
  }
};
</script>
  
<style>
img {
  width: 100%;
  margin: 0 auto;
  display: block;
}
</style>
