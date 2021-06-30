<template>
  <div>
    <Header />
  </div>
  <div class="choose">
    <div class="choosebl">
      <select name="chose" @change="sortBy">
        <option value="sort-asc">Price low to hight</option>
        <option value="sort-desc">Price hight to low</option>
        <option value="sorta-z">Name a to z</option>
        <option value="sortz-a">Name z to a</option>
      </select>
    </div>
    <div class="choosebl">
      <select name="chose" id="chooose2">
        <option value="volvo">1000 - 200000</option>
        <option value="saab">Saab</option>
        <option value="mercedes">Mercedes</option>
        <option value="audi">Audi</option>
      </select>
    </div>
     <div class="card">
      <input type="text" v-model="searchText" @keyup="search()" />
      <button>Search</button>
    </div>
  </div>
  <div class="container">
    <div class="block" v-for="item in motorbikes" :key="item.id">
      <img :src="item.image" alt="" />
      <h3>{{ item.title }}</h3>
      <p>$ {{ item.price }} </p>
    </div>
  </div>
</template>

<script>
import axios from 'axios'
import Header from "./components/Header.vue";

export default {
  name: "App",
  components: { Header },
  data() {
    return {
      motorbikes: [
     
      ],
      msg: "...",
    };
  },
  methods: {
    countProduct() {
      const product = this.motorbikes.filter((motorbike) => motorbike.id);
      return product.length;
    },
    totalPriceOfMotor() {
      let sum = 0;
      this.motorbikes.forEach((motorbike) => {
        sum += motorbike.price;
      });
      return sum;
    },
    avgPriceOfMotor() {
      return (this.totalPriceOfMotor() / this.motorbikes.length).toFixed(2);
    },
    sortBy(event) {
      console.log(event.target.value);
      this.msg = event.target.value;
      if (event.target.value == "sort-asc") {
        this.motorbikes = this.sortByPrice(this.motorbikes, "asc");
      } else if (event.target.value == "sort-desc") {
        this.motorbikes = this.sortByPrice(this.motorbikes, "desc");
      } else if (event.target.value == "sorta-z") {
        this.motorbikes = this.sortByName(this.motorbikes, "asc");
      } else if (event.target.value == "sortz-a") {
        this.motorbikes = this.sortByName(this.motorbikes, "desc");
      }
    },
    sortByPrice(motor, compare) {
      if (compare == "asc") {
        motor.sort(function (a, b) {
          return a.price - b.price;
        });
        return motor;
      } else if (compare == "desc") {
        motor.sort(function (a, b) {
          return -a.price + b.price;
        });
        return motor;
      }
    },
    sortByName(motor, compare) {
      if (compare == "asc") {
        motor.sort(function (a, b) {
          return a.name.localeCompare(b.name);
        });
        return motor;
      } else if (compare == "desc") {
        motor.sort(function (a, b) {
          return b.name.localeCompare(a.name);
        });
        return motor;
      }
    },
    search() {
			this.inItData = false;
			this.productSearchs = [];
			this.motorbikes.forEach((motor) => {
				if (motor.name.toLowerCase().includes(this.searchQuery.toLowerCase())) {
					this.productSearchs.push(motor);
				}
			});
			if (this.searchQuery.length === 0) {
				this.productSearchs = this.motor;
				this.results = false;
    }},
  },
  mounted(){
    axios
      // .get("https://fake-rest-api-nodejs.herokuapp.com/products")
      .get('http://localhost:3000/products')
      .then((response) => {
        this.motorbikes = Object.values(response.data);
        console.log(this.motorbikes)
      });
  }
};
</script>

<style>
.container {
  margin: auto;
  width: 70%;
  display: flex;
  padding: 40px;
  flex-wrap: wrap;
}
img {
  width: 300px;
  height: 210px;
}
.block {
  margin: 10px;
  display: flex;
  flex-direction: column;
  align-items: center;
}
.choose {
  display: flex;
  justify-content: center;
}
.choosebl {
  margin: 60px;
}
/* .card {
  display: flex;
  justify-content: center;
} */
</style>
