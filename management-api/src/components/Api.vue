<template>
  <div class="view">
    <h1>{{msg}}</h1>
    <router-link to="/insertapi" class="insert is-link button is-small is-pulled-right">InsertAPI</router-link>
    <table class="table">
      <thead>
        <tr>
          <th><abbr title="Name">Name</abbr></th>
          <th><abbr title="Path">Path</abbr></th>
          <th><abbr title="Method">Method</abbr></th>
          <th><abbr title="Access Path">Access Path</abbr></th>
          <th><abbr title="Setting">Setting</abbr></th>
        </tr>
      </thead>
      <tfoot>
        <tr>
          <th><abbr title="Name">Name</abbr></th>
          <th><abbr title="Path">Path</abbr></th>
          <th><abbr title="Method">Method</abbr></th>
          <th><abbr title="Access Path">Access Path</abbr></th>
          <th><abbr title="Setting">Setting</abbr></th>
        </tr>
      </tfoot>
      <tbody>
        <tr v-for="(v ,i) in this.list" :key="i">
          <th>{{v.name}}</th>
          <td>{{v.path}}</td>
          <td style="border-right: solid 1px #BDBDBD;">{{v.method}}</td>
          <td>{{v.accessPath}}</td>
          <td><a class="button is-small is-success">Edit</a><a class="button is-small is-danger">Delete</a></td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
export default {
  name: 'Api',
  data () {
    return {
      msg: 'Welcome to Api!!',
      list: []
    }
  },
  created: function () {
    this.index()
  },
  methods: {
    index: function () {
      this.$store.dispatch('api/index', {})
        .then(data => {
          for (let item of data) {
            this.list.push({name: item.name})
          }
        })
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.insert{
  margin: 10px;
}
.button{
  margin: 3px;
}
</style>
