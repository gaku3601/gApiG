<template>
  <div class="sidemenu">
    <aside class="menu">
      <p class="menu-label">
        Menu
      </p>
      <ul class="menu-list">
        <li><a>Dashboard</a></li>
        <li><a>Admin List</a></li>
        <li><a>User List</a></li>
        <li><a>API Controll</a></li>
        <li><a>Settings</a></li>
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
