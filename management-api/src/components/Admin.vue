<template>
  <div class="view">
    <h1>{{msg}}</h1>
  </div>
</template>

<script>
export default {
  name: 'Admin',
  data () {
    return {
      msg: 'Welcome to Admin!!'
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.view{
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
