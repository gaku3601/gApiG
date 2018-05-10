<template>
  <div class="sidemenu">
    <aside class="menu">
      <p class="menu-label">
        Menu
      </p>
      <ul class="menu-list">
        <li><router-link to="/management">Dashboard</router-link></li>
        <li><router-link to="/admin">Admin List</router-link></li>
        <li><router-link to="/user">User List</router-link></li>
        <li><router-link to="/api">Api Controll</router-link></li>
        <li><router-link to="/setting">Setting</router-link></li>
      </ul>
    </aside>
    </div>
  </div>
</template>

<script>
export default {
  name: 'SideMenu',
  data () {
    return {
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.sidemenu{
  border-right: solid 1px #BDBDBD;
}
</style>
