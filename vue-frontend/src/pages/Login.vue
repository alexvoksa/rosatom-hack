<template>
  <div class="form-container">
    <div class="img-container">
      <img src="@/assets/rosatom_logo.png" alt="" class="resize" />
    </div>
    <el-card class="box-card form">
      <h2 slot="header" class="clearfix">Login</h2>
      <div>Please enter your login and password</div>
      <div class="input-container">
        <el-input placeholder="Username" v-model="username"></el-input>
      </div>
      <div class="input-container">
        <el-input placeholder="Password" v-model="password"></el-input>
      </div>
      <div class="button-container">
        <el-button type="primary" @click="login">Login</el-button>
      </div>
    </el-card>
  </div>
</template>

<script>
export default {
  data() {
    return {
      username: null,
      password: null,
    };
  },
  methods: {
    login() {
      
    }
  }
};
</script>

<style scoped>
.form-container {
  padding-top: 50px;
  width: 100vw;
  background-color: rgb(99, 133, 207);
  height: 100vh;
}

.img-container {
  padding: 50px;
}

img.resize {
  width: 400px; /* you can use % */
  height: auto;
}

.form {
  background-color: whitesmoke;
  display: block;
  margin-left: auto;
  margin-right: auto;
  width: 30%;
  display: flex;
  flex-direction: column;
}

.input-container {
  padding-left: 50px;
  padding-right: 50px;

  margin-top: 50px;
}

.button-container {
  margin-top: 64px;
}
</style>