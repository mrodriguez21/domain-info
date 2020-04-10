<template>
  <div class="container">
    <h1 class="header">Previously consulted domains</h1>
    <ul>
      <div class="domain-collapsed" v-for="item in items" :key="item.domain">
        <Domain collapsed :name="item.domain" :domain="item.info"></Domain>
      </div>
    </ul>
  </div>
</template>

<script>
import Domain from "@/components/Domain.vue";
import APIClient from "../apiClient";

export default {
  name: "History",
  components: {
    Domain,
  },
  data() {
    return {
      items: [],
    };
  },
  created() {
    this.getHistory();
  },
  methods: {
    getHistory() {
      APIClient.getPreviousDomains().then((response) => {
        this.items = response.items;
      });
    },
  },
};
</script>

<style>
.domain-collapsed {
  margin-top: 20px;
  margin-bottom: 35px;
}
.header {
  font-size: 2rem;
}
</style>
