<template>
  <div>
    <v-row align="center">
      <v-col class="text-center" cols="12" sm="2">
        <div class="my-2">
          <v-btn text>日付順</v-btn>
        </div>
      </v-col>

      <v-col class="text-center" cols="12" sm="2">
        <div class="my-2">
          <v-btn text>いいね順</v-btn>
        </div>
      </v-col>

      <v-col class="text-center" cols="12" sm="2">
        <div class="my-2">
          <v-btn text>回答数順</v-btn>
        </div>
      </v-col>
    </v-row>
  <v-divider />
  <br>
  </div>
</template>


<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator';

@Component
export default class QuestionSetting extends Vue {
}
</script>

<style>
</style>
