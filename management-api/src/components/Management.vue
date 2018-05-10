<template>
  <div class="Management">
    <div class="container">
      <div class="header">Header {{msg}}</div>
      <div class="sidemenu"><SideMenu /></div>
      <div class="content">
        <router-view></router-view>
      </div>
      <div class="footer">footer</div>
    </div>
  </div>
</template>

<script>
import SideMenu from './SideMenu'
export default {
  name: 'Management',
  data () {
    return {
      msg: 'Welcome to Management!! is Header'
    }
  },
  components: {
    SideMenu
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.container{
  min-height: 100vh;
  display: grid;
  grid-template-rows: 100px 1fr 30px;
  grid-template-columns: 150px 1fr;
}
.header{
  grid-row: 1 / 2;
  grid-column: 1 / 3;
}
.sidemenu{
  grid-row: 2 / 3;
  grid-column: 1 / 2;
}
.content{
  grid-row: 2 / 3;
  grid-column: 2 / 3;
}
.footer{
  grid-row: 3 / 4;
  grid-column: 1 / 3;
  padding: 0;
}
</style>
