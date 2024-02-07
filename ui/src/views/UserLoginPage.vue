<script setup lang="ts">
import { ref } from "vue";
import { useRouter } from "vue-router";
import { userLogin } from "../api/user";

const router = useRouter();

const username = ref("");
const password = ref("");

const login = async () => {
  const resp = await userLogin(username.value, password.value);
  if (resp.code === 200) {
    localStorage.setItem("accessToken", resp.data.token);
    router.push({ name: "card-page" });
    return;
  }
};
</script>

<template>
  <section>
    <div class="box">
      <div class="square wapnone" style="--i: 0"></div>
      <div class="square wapnone" style="--i: 1"></div>
      <div class="square wapnone" style="--i: 2"></div>
      <div class="square wapnone" style="--i: 3"></div>
      <div class="square wapnone" style="--i: 4"></div>
      <div class="square wapnone" style="--i: 5"></div>

      <div class="container">
        <div class="login-form">
          <h2>日拱一卒</h2>
          <div class="inputBx">
            <input type="text" required="required" v-model="username" />
            <span>用户名</span>
            <i class="fas"></i>
          </div>
          <div class="inputBx password">
            <input type="password" v-model="password" required="required" @keyup.enter.native="login()" />
            <span>密码</span>
            <i class="fas"></i>
          </div>
          <div class="inputBx">
            <input type="submit" value="登录" @click="login()" />
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<style scoped lang="scss">
.fas {
  width: 32px;
}

@media screen and (max-width: 768px) {
  .wapnone {
    display: none;
  }
}

section {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 100vh;
  width: 100%;
  background: linear-gradient(-45deg, #ee7752, #e73c7e, #23a6d5, #23d5ab);
  background-size: 400% 400%;
}

.box {
  position: relative;

  .square {
    position: absolute;
    background: rgba(255, 255, 255, 0.1);
    backdrop-filter: blur(5px);
    box-shadow: 0 25px 45px rgba(0, 0, 0, 0.1);
    border: 1px solid rgba(255, 255, 255, 0.15);
    border-radius: 15px;
    animation: square 10s linear infinite;
    animation-delay: calc(-1s * var(--i));
  }

  @keyframes square {

    0%,
    100% {
      transform: translateY(-20px);
    }

    50% {
      transform: translateY(20px);
    }
  }

  .square:nth-child(1) {
    width: 100px;
    height: 100px;
    top: -15px;
    right: -45px;
  }

  .square:nth-child(2) {
    width: 150px;
    height: 150px;
    top: 105px;
    left: -125px;
    z-index: 2;
  }

  .square:nth-child(3) {
    width: 60px;
    height: 60px;
    bottom: 85px;
    right: -45px;
    z-index: 2;
  }

  .square:nth-child(4) {
    width: 50px;
    height: 50px;
    bottom: 35px;
    left: -95px;
  }

  .square:nth-child(5) {
    width: 50px;
    height: 50px;
    top: -15px;
    left: -25px;
  }

  .square:nth-child(6) {
    width: 85px;
    height: 85px;
    top: 165px;
    right: -155px;
    z-index: 2;
  }
}

.container {
  position: relative;
  padding: 50px;
  width: 260px;
  min-height: 380px;
  display: flex;
  justify-content: center;
  align-items: center;
  background: rgba(255, 255, 255, 0.1);
  backdrop-filter: blur(5px);
  border-radius: 10px;
  box-shadow: 0 25px 45px rgba(0, 0, 0, 0.2);
}

.container::after {
  content: "";
  position: absolute;
  top: 5px;
  right: 5px;
  bottom: 5px;
  left: 5px;
  border-radius: 5px;
  pointer-events: none;
  background: linear-gradient(to bottom,
      rgba(255, 255, 255, 0.1) 0%,
      rgba(255, 255, 255, 0.1) 2%);
}

.login-form {
  position: relative;
  width: 100%;
  height: 100%;

  h2 {
    color: #fff;
    letter-spacing: 2px;
    margin-bottom: 30px;
  }

  .inputBx {
    position: relative;
    width: 100%;
    margin-bottom: 20px;

    input {
      width: 80%;
      outline: none;
      border: none;
      border: 1px solid rgba(255, 255, 255, 0.2);
      background: rgba(255, 255, 255, 0.2);
      padding: 8px 10px;
      padding-left: 40px;
      border-radius: 15px;
      color: #fff;
      font-size: 16px;
      box-shadow: 0 5px 15px rgba(0, 0, 0, 0.05);
    }

    .fas {
      position: absolute;
      top: 13px;
      left: 13px;
    }

    input[type="submit"] {
      background: #fff;
      color: #111;
      max-width: 100px;
      padding: 8px 10px;
      box-shadow: none;
      letter-spacing: 1px;
      cursor: pointer;
      transition: 1.5s;
    }

    input::placeholder {
      color: #fff;
    }

    span {
      position: absolute;
      left: 30px;
      padding: 10px;
      display: inline-block;
      color: #fff;
      transition: 0.5s;
      pointer-events: none;
    }

    input:focus~span,
    input:valid~span {
      transform: translateX(-30px) translateY(-25px);
      font-size: 12px;
    }
  }
}
</style>
