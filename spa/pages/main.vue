<template>
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
</template>

<script>
import axios from "axios";
export default {
  data() {
    return {
      foo: "hoge",
      jobs: [],
      removeId: null,
      isAdd: false,
      test: "#abc"
    };
  },
  mounted() {
    axios.get("https://hal-iot.net/api/admin/jobs").then(res => {
      this.jobs = res.data.jobs
        .reverse()
        .slice(0, 7)
        .reverse();
      // this.jobs.reverse();
      // this.jobs = this.jobs.slice(0, 5);
    });
    setInterval(this.change, 500);
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
          }, 1000);
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
  background-image: url("~assets/main_back.png");
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
  font-family: "M PLUS Rounded 1c";
  font-size: 1.5rem;
}

.task-box {
  list-style: none;
  width: 68vw;
  height: 30vw;
  border: 1px solid dimgrey;
  border-radius: 7vw;
}
.job-name {
  font-size: 2rem;
  text-align: center;
  margin-top: 7vw;
}
.user-name {
  font-size: 1rem;
  text-align: right;
  margin-top: 3vw;
  margin-right: 3vw;
}

/* アニメーション */
.task-box.remove {
  animation-name: slide_out;
  animation-duration: 1s;
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

