<template>
  <div class="view">
    <h1>{{msg}}</h1>
  </div>
</template>

<script>
export default {
  name: 'Api',
  data () {
    return {
      msg: 'Welcome to Api!!'
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
