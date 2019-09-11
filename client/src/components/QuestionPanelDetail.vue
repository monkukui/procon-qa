<template>
  <div class="question-panel-detail">

    <v-card
      max-width="1010"
      class="mx-auto"
      color="white"
    >
      <v-card-title>
        {{ title }}
      </v-card-title>
      <v-btn
        v-for="tag in tags"  
        class="tag"
        color="blue-grey lighten-4"
        x-small
      >
        {{ tag }}
      </v-btn>
      
      <v-divider class="mx-4"></v-divider>
      <v-card-text>{{ body }}</v-card-text>
      <v-card-text>質問日時: {{ questionedTime }}</v-card-text>
    </v-card>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue, Emit } from 'vue-property-decorator';

@Component
export default class QuestionPanel extends Vue {
  @Prop()
  private title!: string;
  @Prop()
  private body!: string;
  @Prop()
  private questionedTime!: string;
  @Prop()
  private answered!: boolean;
  private tags: string[] = [
    'AtCoder',
    'ABC',
    'C++',
  ];
  private openQuestionPage(): void {
    this.$router.push({
      path: '/question',
      query: {
        title: this.title,
        body: this.body,
        questionedTime: this.questionedTime,
      },
    });
  }

}
</script>

<style scoped>
.small {
  font-size: 75%;
}
.question-panel-detail {
  margin: 3%;
}

.tag {
  margin-bottom: 1%;
  margin-left: 2%;
}
</style>
