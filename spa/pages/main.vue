<template>
  <div>
    <ul id="v-for-object" class="task-boxs">
      <li
        v-for="(value, index) in jobs"
        :key="index"
        class="task-box"
        :class="{ remove: value.id === removeId, add: index === 0 && isAdd }"
        :style="{'background-color':value.image_color}"
      >
        <p class="job-name">{{ value.display_name }}</p>
        <p class="user-name">{{ value.user_name }}</p>
      </li>
    </ul>
    <audio ref="audio">
      <source ref="source" :src="se" type="audio/mpeg">Your browser does not support the audio element.
    </audio>
  </div>
</template>

<script>
import axios from "axios";
import io from "socket.io-client";
export default {
  data() {
    return {
      foo: "hoge",
      jobs: [],
      removeId: null,
      isAdd: false,
      test: "#abc",
      se: ""
    };
  },
  mounted() {
    axios.get("https://hal-iot.net/api/admin/jobs").then(res => {
      this.jobs = res.data.jobs
        .reverse()
        .slice(0, 7)
        .reverse();
    });
    setInterval(this.change, 500);
    // 効果音
    this.socket = io("https://socket.patra.store", {
      transports: ["websocket"]
    });
    this.socket.on("se", msg => {
      this.se = msg;
      this.$refs.audio.load();
      this.$refs.audio.play();
    });
  },
  methods: {
    change() {
      axios.get("https://hal-iot.net/api/admin/jobs").then(res => {
        var newJobs = res.data.jobs
          .reverse()
          .slice(0, 7)
          .reverse();
        if (
          this.jobs[this.jobs.length - 1].id !== newJobs[newJobs.length - 1].id
        ) {
          this.removeId = this.jobs[this.jobs.length - 1].id;
          setTimeout(() => {
            this.jobs = newJobs;
          }, 350);
        }
        if (this.jobs.length < newJobs.length) {
          this.jobs = newJobs;
          this.isAdd = true;
        } else {
          this.isAdd = false;
        }
      });
    }
  }
};
</script>

<style>
body,
html {
  width: 100%;
  height: 100%;
}

body {
  background-image: url("~assets/main_back.svg");
  background-repeat: no-repeat;
  background-size: 100%;
  background-position: 0 0;
}

*,
*::before,
*::after {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

.task-boxs {
  height: 100vh;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: flex-end;
  font-family: "Kosugi Maru", sans-serif;
  font-size: 1.5rem;
}

.task-box {
  list-style: none;
  width: 68vw;
  height: 14vh;
  border: 1px solid dimgrey;
  border-radius: 7vw;
  opacity: 0.95;
  box-shadow: 20px 20px 40px;
}
.job-name {
  font-size: 5rem;
  text-align: center;
  margin-top: 4vh;
  font-weight: bold;
}
.user-name {
  font-size: 3rem;
  text-align: right;
  margin-top: 2.5vw;
  margin-right: 3vw;
}

/* アニメーション */
.task-box.remove {
  animation-name: slide_out;
  animation-duration: 0.6s;
}

@keyframes slide_out {
  0% {
  }
  100% {
    transform: translateX(100vw);
  }
}

.task-box.add {
  animation: poyon 1.1s linear 0s 1;
}

@keyframes poyon {
  0% {
    transform: scale(0.8, 1.4) translate(0%, -100%);
  }
  10% {
    transform: scale(0.8, 1.4) translate(0%, -15%);
  }
  20% {
    transform: scale(1.4, 0.6) translate(0%, 30%);
  }
  30% {
    transform: scale(0.9, 1.1) translate(0%, -10%);
  }
  40% {
    transform: scale(0.95, 1.2) translate(0%, -30%);
  }
  50% {
    transform: scale(0.95, 1.2) translate(0%, -10%);
  }
  60% {
    transform: scale(1.1, 0.9) translate(0%, 5%);
  }
  70% {
    transform: scale(1, 1) translate(0%, 0%);
  }
  100% {
    transform: scale(1, 1) translate(0%, 0%);
  }
}
</style>

