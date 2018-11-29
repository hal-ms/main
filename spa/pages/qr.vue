<template>
  <div>
    <div id="container">
      <div id="roop"></div>
      <div id="title-container">
        <div id="title">お仕事の依頼はこちらから！</div>
        <div id="logo">
          <img src="~assets/logo.png" alt="三日坊主のロゴ">
        </div>
      </div>
      <div id="truck-container">
        <div id="truck" :class="{ move: isMove }">
          <img src="~assets/truck.png" alt="トラックの画像">
          <div id="qr-code-container">
            <qrcode-vue :value="url" size="180" level="H"></qrcode-vue>
          </div>
        </div>
      </div>
    </div>
    <audio>
      <source src="https://hal-iot.net/public/vanilla.mp3" type="audio/mpeg">Your browser does not support the audio element.
    </audio>
  </div>
</template>

<script>
import io from "socket.io-client";
import QrcodeVue from "qrcode.vue";
import axios from "axios";

export default {
  data() {
    return {
      url: "",
      isMove: false
    };
  },
  components: {
    QrcodeVue
  },
  mounted() {
    this.GetToken();
    this.socket = io("https://socket.patra.store", {
      transports: ["websocket"]
    });
    this.socket.on("qr", msg => {
      this.GetToken();
    });
  },
  methods: {
    GetToken: function() {
      axios.get("https://hal-iot.net/api/token").then(res => {
        this.isMove = true;
        console.log("move");
        setTimeout(() => {
          this.url = "https://hal-iot.net/create/" + res.data.id;
          console.log(res.data.id);
          console.log("move end");
          document.querySelector("audio").play();
          this.isMove = false;
        }, 3000);
      });
    }
  }
};
</script>

<style>
audio {
  height: 30px;
}
#container {
  position: relative;
  width: 100%;
  height: 100%;
}
#roop {
  width: 100%;
  height: 600px;
  background: url("~assets/background.png") repeat-x;
  background-position: 0 0;
  background-size: auto 600px;
  -webkit-animation: bgroop 20s linear infinite;
  animation: bgroop 20s linear infinite;
}
@-webkit-keyframes bgroop {
  from {
    background-position: 0 0;
  }
  to {
    background-position: 1839px 0;
  }
}
@keyframes bgroop {
  from {
    background-position: 0 0;
  }
  to {
    background-position: 1839px 0;
  }
}
#truck-container {
  -webkit-animation: fluffy 2s ease infinite;
  animation: fluffy 2s ease infinite;
}
#truck {
  position: absolute;
  right: 30px;
  bottom: 27px;
  width: 550px;
  opacity: 0;
  -webkit-animation: to_start_position 1s linear;
  animation: to_start_position 1s linear;
  animation-fill-mode: forwards;
}
@keyframes fluffy {
  0% {
    transform: translateY(0);
  }
  5% {
    transform: translateY(0);
  }
  10% {
    transform: translateY(-5px);
  }
  20% {
    transform: translateY(0);
  }
  25% {
    transform: translateY(0);
  }
  30% {
    transform: translateY(-10px);
  }
  50% {
    transform: translateY(0);
  }
  100% {
    transform: translateY(0);
  }
}
#truck img {
  display: block;
  width: 100%;
}
#truck.move {
  -webkit-animation: acceleration 2s linear;
  animation: acceleration 2s linear;
}
@keyframes acceleration {
  0% {
    opacity: 1;
    right: 30px;
  }
  99% {
    opacity: 1;
  }
  100% {
    opacity: 0;
    right: 100%;
  }
}
@keyframes to_start_position {
  from {
    opacity: 1;
    right: -600px;
  }
  to {
    opacity: 1;
    right: 30px;
  }
}

#qr-code-container {
  position: absolute;
  right: 100px;
  bottom: 80px;
  width: 200px;
  height: 200px;
  padding: 10px;
  background-color: white;
}

#title-container {
  position: absolute;
  top: 30px;
  left: 0;
  padding: 20px;
}
#title {
  font-size: 50px;
  color: #114f97;
}
#logo {
  width: 350px;
}
#logo img {
  display: block;
  width: 100%;
}
</style>

