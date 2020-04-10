<template>
  <div class="container">
    <div id="datainfo">
      <div id="status">
        <mdb-tooltip trigger="hover" :options="{ placement: 'left' }">
          <span slot="tip">
            Current SSL grade
          </span>
          <div slot="reference" class="ssl_grade mb-1">
            <i class="ssl_icon fas fa-circle fa-2x" :class="iconColor"></i>
            <span :class="textColor">{{ domain.ssl_grade }}</span>
          </div>
        </mdb-tooltip>
        <mdb-tooltip trigger="hover" :options="{ placement: 'left' }">
          <span slot="tip">
            Previous SSL grade
          </span>
          <div slot="reference" class="ssl_grade">
            <i class="ssl_icon fas fa-circle fa-2x" :class="prevIconColor"></i>
            <span :class="prevTextColor">
              {{
                domain.previous_ssl_grade == ""
                  ? "-"
                  : domain.previous_ssl_grade
              }}
            </span>
          </div>
        </mdb-tooltip>
      </div>
      <div id="fav">
        <img
          class="domain-logo"
          :src="domain.logo == '' ? '../assets/not-found.png' : domain.logo"
          alt="Logo"
        />
      </div>
      <div id="info">
        <div id="subinfo">
          <span class="domain-name title">{{ name }}</span>
          <div>
            <mdb-tooltip trigger="hover" :options="{ placement: 'right' }">
              <span slot="tip">
                Host is currently
                <b>{{ domain.is_down ? "down" : "up" }}.</b>
              </span>
              <i
                slot="reference"
                class="fas ml-2 fa-2x"
                :class="[
                  domain.is_down
                    ? 'fa-times-circle red-text'
                    : 'fa-check-circle green-text',
                ]"
              ></i>
            </mdb-tooltip>
          </div>
        </div>
        <div>{{ domain.title }}</div>
      </div>
    </div>
    <div class="title">Servers</div>
    <b-table striped :items="domain.servers" :fields="fields">
      <template v-slot:table-caption>
        <p class="text-center">
          *Servers have {{ domain.servers_changed ? "" : "not" }} changed since
          last request received.
        </p>
      </template>
    </b-table>
  </div>
</template>

<script>
import { mdbTooltip } from "mdbvue";
export default {
  components: { mdbTooltip },
  computed: {
    iconColor() {
      switch (this.domain.ssl_grade) {
        case "A+":
          return "green-text";
        case "A":
          return "light-green-text";
        case "A-":
          return "lime-text";
        case "B":
          return "yellow-text";
        case "C":
          return "amber-text";
        case "D":
          return "orange-text";
        case "E":
          return "deep-orange-text";
        case "F":
          return "red-text";
        default:
          return "text-light";
      }
    },
    prevIconColor() {
      switch (this.domain.previous_ssl_grade) {
        case "A+":
          return "green-text";
        case "A":
          return "light-green-text";
        case "A-":
          return "lime-text";
        case "B":
          return "yellow-text";
        case "C":
          return "amber-text";
        case "D":
          return "orange-text";
        case "E":
          return "deep-orange-text";
        case "F":
          return "red-text";
        default:
          return "text-light";
      }
    },
    textColor() {
      return this.domain.ssl_grade == "B" ? "text-black-50" : "white-text";
    },
    prevTextColor() {
      return this.domain.previous_ssl_grade == "B"
        ? "text-black-50"
        : "white-text";
    },
  },
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
    name: String,
    domain: Object,
  },
};
</script>

<style>
#datainfo {
  display: flex;
  flex: 1;
  flex-direction: row;
  justify-items: center;
  align-items: stretch;
  width: 60vw;
  margin: auto;
  margin-bottom: 16px;
  border: 1px solid lightgray;
  padding: 20px;
  border-radius: 10px;
  /*! box-shadow: 0 1px 13px -7px black; */
}
#status {
  flex-direction: column;
  display: flex;
  justify-content: space-around;
  margin-right: 10px;
}
#info {
  display: flex;
  flex-direction: column;
  flex: 5;
  flex-grow: 1;
}
#subinfo {
  flex-direction: row;
  display: flex;
  flex-grow: 1;
  align-items: center;
  justify-items: self-end;
  align-self: center;
}
#subinfo > div {
  justify-self: center;
  align-self: center;
}
.ssl_grade {
  position: relative;
  display: flex;
  height: 2em;
  width: 2em;
}
.ssl_grade > i {
  position: absolute;
}
.ssl_grade > span {
  margin: auto;
  z-index: 10;
  font-weight: bold;
}

.domain-logo {
  height: 4.25rem;
}
.domain-name {
  padding-bottom: 5px;
}

.title {
  font-size: 1.5rem;
}

thead tr th {
  font-weight: bold;
}
</style>
