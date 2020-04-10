<template>
  <div>
    <div class="title">Servers</div>
    <b-table striped :items="servers" :fields="fields">
      <template v-slot:table-caption>
        <p class="text-center">
          *Servers have {{ servers_changed ? "" : "not" }} changed since last
          request received.
        </p>
      </template>
    </b-table>
  </div>
</template>

<script>
export default {
  data() {
    return {
      fields: [
        {
          key: "address",
          sortable: true,
        },
        {
          key: "ssl_grade",
          label: "SSL Grade",
          sortable: true,
        },
        {
          key: "country",
          sortable: true,
        },
        {
          key: "owner",
          sortable: true,
        },
      ],
    };
  },
  props: {
    servers: Array,
    servers_changed: Boolean,
  },
};
</script>

<style>
thead tr th {
  font-weight: bold;
}
</style>
