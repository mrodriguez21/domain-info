<template>
  <mdb-footer color="indigo" class="font-small pt-0">
    <mdb-container class="text-left">
      <mdb-row class="py-4 d-flex align-items-center">
        <mdb-col md="6" lg="5" class="text-center text-md-left mb-4 mb-md-0">
          <h6 class="mb-0 text-light">
            Developed with ❤️ by
            <a href="https://mrodriguez21.github.io/">Mariana Rodriguez</a>
            &copy; {{ new Date().getFullYear() }}
          </h6>
        </mdb-col>
        <mdb-col md="6" lg="7" class="text-center text-md-right">
          <a class="gi-ic">
            <i class="fab fa-github white-text mr-lg-4"></i>
          </a>
          <a class="li-ic">
            <i class="fab fa-linkedin-in white-text mr-lg-4"></i>
          </a>
          <a class="fb-ic ml-0">
            <i class="fab fa-facebook white-text mr-lg-4"></i>
          </a>
          <a class="tw-ic">
            <i class="fab fa-twitter white-text mr-lg-4"></i>
          </a>
          <a class="ins-ic">
            <i class="fab fa-instagram white-text mr-lg-4"></i>
          </a>
        </mdb-col>
      </mdb-row>
    </mdb-container>
  </mdb-footer>
  <!-- Footer -->
</template>

<script>
import { mdbFooter, mdbContainer, mdbRow, mdbCol } from "mdbvue";
export default {
  name: "FooterPage",
  components: {
    mdbFooter,
    mdbContainer,
    mdbRow,
    mdbCol
  }
};
</script>
