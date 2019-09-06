<template>
  <div class="question-panel-detail">
    <div class="box">
      <h4>{{ title }}</h4>
      <hr>
      <p>{{ body }}</p>
      <span class="small">質問日時: {{ questionedTime }}</span>
    </div>
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
  font-size: 75%;
}

.box{
  width: 75%;
  padding: 0.5em 1em;
  margin: 2em 0;
  color: #5d627b;
  background: white;
  border-top: solid 10px #5d627b;
  box-shadow: 0 3px 5px rgba(0, 0, 0, 0.22);
}
.box p {
  margin: 0; 
  padding: 0;
}
  
</style>
