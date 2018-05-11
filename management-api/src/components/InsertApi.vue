<template>
  <div class="view">
    <h1>{{msg}}</h1>
    <div class="field">
      <label class="label">Name</label>
      <div class="control has-icons-left">
        <input class="input" type="text" placeholder="Input API Name">
        <span class="icon is-small is-left">
          <i class="fas fa-user"></i>
        </span>
      </div>
    </div>
    <div class="field">
      <label class="label">Path</label>
      <div class="control has-icons-left">
        <input class="input" type="text" placeholder="Input Path">
        <span class="icon is-small is-left">
          <i class="fas fa-user"></i>
        </span>
      </div>
    </div>
    <div class="field">
      <label class="label">Method</label>
      <div class="control has-icons-left">
        <div class="select">
          <select>
            <option>GET</option>
            <option>POST</option>
            <option>PUT</option>
            <option>DELETE</option>
          </select>
        </div>
        <div class="icon is-small is-left">
          <i class="fas fa-globe"></i>
        </div>
      </div>
    </div>
    <div class="field">
      <label class="label">Access Path</label>
      <div class="control has-icons-left">
        <input class="input" type="text" placeholder="Input Access Path">
        <span class="icon is-small is-left">
          <i class="fas fa-user"></i>
        </span>
      </div>
    </div>
    <router-link to="/api" class="is-link button is-small is-pulled-right">Regist</router-link>
    <router-link to="/api" class="is-danger button is-small is-pulled-right">Cancel</router-link>
  </div>
</template>

<script>
export default {
  name: 'Api',
  data () {
    return {
      msg: 'Welcome to ApiInsert!!',
      data: [
        {name: 'Delete Image', path: '/image/:id', method: 'DELETE', accessPath: 'http://imagecontainer:8080/'}
      ]
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.button{
  margin: 10px;
}
</style>
