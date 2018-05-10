<template>
  <div class="login">
    <h1>{{msg}}</h1>
    <div class="form">
      <div class="field">
        <div class="control">
          <input class="input" type="text" placeholder="Your Email">
        </div>
      </div>
      <div class="field">
        <div class="control">
          <input class="input" type="text" placeholder="Password">
        </div>
      </div>
      <div class="control">
        <router-link to="/management">
          <button class="button is-primary">Login</button>
        </router-link>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Login',
  data () {
    return {
      msg: 'Welcome to gApiG!!'
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.login{
  margin-top:60px;
  text-align: center;
}
.form{
  width: 300px;
  margin: 0 auto;
}
button{
  width: 100%;
}
</style>
