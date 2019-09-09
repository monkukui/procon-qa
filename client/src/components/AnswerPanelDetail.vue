<template>
  <div class="answer-panel-detail">
    <v-card
      max-width="1010"
      class="mx-auto"
      color="white"
    >
      <v-card-title>
        回答
      </v-card-title>
      
      <v-divider class="mx-4"></v-divider>
      <v-card-text>{{ body }}</v-card-text>
      <v-card-text>回答日時: {{ questionedTime }}</v-card-text>
    </v-card>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue, Emit } from 'vue-property-decorator';

@Component
export default class AnswerPanelDetail extends Vue {
  @Prop()
  private body!: string;
  @Prop()
  private answeredTime!: string;
}
</script>

<style scoped>

.small {
  font-size: 75%;
}

.answer-panel-detail {
  margin: 2%;
}
  
</style>
