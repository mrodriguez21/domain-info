<template>
  <mdb-navbar position="top" expand="medium" color="indigo" dark>
    <mdb-navbar-brand href="/">
      <a class="text-reset" href="/">
        <img src="../assets/images/logo.png" height="30" alt="" />
        Domain info
      </a>
    </mdb-navbar-brand>
    <!-- Collapse button -->
    <mdb-navbar-toggler>
      <mdb-navbar-nav right>
        <mdb-nav-item href="/">Search</mdb-nav-item>
        <mdb-nav-item href="/history">History</mdb-nav-item>
      </mdb-navbar-nav>
    </mdb-navbar-toggler>
  </mdb-navbar>
</template>

<script>
import {
  mdbNavbar,
  mdbNavbarBrand,
  mdbNavbarToggler,
  mdbNavbarNav,
  mdbNavItem,
} from "mdbvue";
export default {
  name: "Navbar",
  components: {
    mdbNavbar,
    mdbNavbarBrand,
    mdbNavbarToggler,
    mdbNavbarNav,
    mdbNavItem,
  },
};
</script>
