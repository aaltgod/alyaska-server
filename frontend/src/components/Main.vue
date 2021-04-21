<template>
  <div class="hello">
    <h1>{{ msg }}</h1>
    <p>
      The Frame will be here
    </p>
    <h3>FILES</h3>
    <ul>
      <li><a href="/" target="_blank" rel="noopener">create a file</a></li>
      <li><a href="/" target="_blank" rel="noopener">open a file</a></li>
    </ul>
    <h3>INFO</h3>
  </div>
</template>

<script>
export default {
  name: 'ALYASKA SERVER',
  props: {
    msg: String
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
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
