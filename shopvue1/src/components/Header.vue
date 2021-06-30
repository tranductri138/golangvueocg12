<template>
  <header class="header">
    <div>
      <p>App vue hihihiihihi</p>
    </div>
    <!-- <div class="card">
      <input type="text" v-model="searchText" @keyup="search()" />
      <button>Search</button>
    </div> -->
    <div class="login">
      <button>Login</button>
    </div>
    <div class="cart">
      <button>
      <img src="@/assets/giohang1.png" alt="">
      </button>
    </div>
  </header>
</template>

<script>
export default {
  name: "Header",
};
</script>

<style>
.header {
  height: 56px;
  padding: 0 16px;
  display: flex;
  flex-direction: row;
  background-color: #ee6363;
  align-items: center;
  justify-content: space-between;
}
.login button {
  height: 30px;
  width: 50px;
  margin: 10px;
}
.card button {
  background-color: #00FF00
}
.card input {
  width: 200px;
}
.cart img{
  height: auto;
  max-width: 40px;
  display: flex; 
}
.cart{
  margin-right : 10px;
}
</style>