<template>
  <div class="question-panel">
    <div v-if="answered" class="box-ac">
      <h4><a class="title" @click="openQuestionPage" href="javascript:void(0);">{{ title }}</a></h4>
      <hr>
      <p>{{ body }}</p>
      <span class="small">質問日時: {{ questionedTime }}</span>
    </div>
    <div v-else class="box-wa">
      <h4><a class="title" @click="openQuestionPage" href="javascript:void(0);">{{ title }}</a></h4>
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

a {
  color: #005999;
  text-decoration: none;
}
a:hover {
  color: #0077cc;
  text-decoration: none;
}

.small {
  font-size: 75%;
}

.question-panel {
}
.box-ac{
  width: 75%;
  padding: 0.5em 1em;
  margin: 2em 0;
  color: #5d627b;
  background: white;
  border-top: solid 10px #5cb85c;
  box-shadow: 0 3px 5px rgba(0, 0, 0, 0.22);
}
.box-ac p {
  margin: 0; 
  padding: 0;
}
  

.box-wa {
  width: 75%;
  padding: 0.5em 1em;
  margin: 2em 0;
  color: #5d627b;
  background: white;
  border-top: solid 10px #f0ad4e;
  box-shadow: 0 3px 5px rgba(0, 0, 0, 0.22);
}
.box-wa p {
  margin: 0; 
  padding: 0;
}
  
</style>
