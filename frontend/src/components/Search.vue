<template>
  <div class="container">
    <Searchbar @search-submitted="getDomain($event)"></Searchbar>
    <Domain v-if="name != ''" :name="name" :domain="domainInfo"></Domain>
  </div>
</template>

<script>
import Searchbar from "@/components/Searchbar.vue";
import Domain from "@/components/Domain.vue";
import APIClient from "../apiClient";

export default {
  name: "App",
  components: {
    Searchbar,
    Domain,
  },
  data() {
    return {
      name: "",
      domainInfo: {},
    };
  },
  methods: {
    getDomain(domainName) {
      this.name = "";
      this.domain = {};
      APIClient.getDomain(domainName).then((response) => {
        this.domainInfo = response;
        this.name = domainName;
      });
    },
  },
};
</script>
