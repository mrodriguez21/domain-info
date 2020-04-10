<template>
  <div class="container">
    <Searchbar @search-submitted="getDomain($event)"></Searchbar>
    <Domain v-if="name != ''" :name="name" :domain="domainInfo"></Domain>
    <Loader v-if="loading"></Loader>
    <Error v-if="error"></Error>
  </div>
</template>

<script>
import Searchbar from "@/components/Searchbar.vue";
import Domain from "@/components/Domain.vue";
import Loader from "@/components/Loader.vue";
import Error from "@/components/Error.vue";
import APIClient from "../apiClient";

export default {
  name: "App",
  components: {
    Searchbar,
    Domain,
    Loader,
    Error,
  },
  data() {
    return {
      name: "",
      domainInfo: {},
      loading: false,
      error: false,
    };
  },
  methods: {
    getDomain(domainName) {
      this.error = false;
      this.name = "";
      this.domainInfo = {};
      this.loading = true;
      APIClient.getDomain(domainName)
        .then((response) => {
          this.loading = false;
          this.domainInfo = response;
          this.name = domainName;
        })
        .catch(() => {
          this.loading = false;
          this.error = true;
        });
    },
  },
};
</script>
