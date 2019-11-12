<template>
  <div class="home">
    <div class="main-contents">
      <v-container fluid>
        <v-row>
          <v-col col="12" sm="8">
            <Questions />
          </v-col>
        </v-row>
      </v-container>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import Questions from '@/components/Questions.vue';

@Component({
  components: {
    Questions,
  },
})
export default class Home extends Vue {
}
</script>

<style scoped>
.home {
}
</style>
