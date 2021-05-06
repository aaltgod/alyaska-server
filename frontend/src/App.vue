<template>
  <div id="app">
    <h1>Alyaska Server</h1>
    <ul>
      <li><router-link to='/'><h2>Home</h2></router-link></li>
      <li><router-link to='/archive'><h2>Archive</h2></router-link></li>
      <li><router-link to='/upload'><h2>Upload</h2></router-link></li>
    <router-view></router-view>
    </ul>
  </div>
</template>

<script>

export default {
} 
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>

<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>

