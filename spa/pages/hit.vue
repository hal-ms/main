<template>
  <div>
    <img :src="img">
    <audio ref="audio2" loop>
      <source ref="source" :src="bgm" type="audio/mpeg">Your browser does not support the audio element.
    </audio>
  </div>
</template>

<script>
import io from "socket.io-client";
export default {
  data() {
    return {
      img:
        "https://hal-iot.net/public/hit/no_wear/standby.jpg",
      bgm: ""
    };
  },
  mounted() {
    this.socket = io("https://socket.patra.store", {
      transports: ["websocket"]
    });
    this.socket.on("hit", msg => {
      this.img = JSON.parse(msg).img;
    });
    this.socket.on("bgm", msg => {
      console.log(msg);
      this.bgm = msg;
      this.$refs.audio2.load();
      this.$refs.audio2.play();
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
