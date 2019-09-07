<template>
  <div class="question-panel">
    <v-card
      max-width="1010"
      class="mx-auto"
      color="white"
    >
      <v-card-title>
        <a class="title" @click="openQuestionPage">{{ title }}</a>
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
      <v-card-text>質問日時: {{ questionedTime }}</v-card-text>
    </v-card>

    <!--div v-if="answered" class="box-ac">
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
    </div-->
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
.question-panel {
  margin: 3%;
}

.tag {
  margin-bottom: 1%;
  margin-left: 2%;
}

.title {
  color: #0288D1;
  text-decoration: none;
}
.title:hover {
  color: #29B6F6;
  text-decoration: none;
}

.small {
  font-size: 75%;
}

.question-panel {
}
.box-ac{
  padding: 0.5em 1em;
  margin: 2em 0;
  color: #5d627b;
  background: white;
  border-top: solid 5px #5cb85c;
  box-shadow: 0 3px 5px rgba(0, 0, 0, 0.22);
}
.box-ac p {
  margin: 0; 
  padding: 0;
}
  

.box-wa {
  padding: 0.5em 1em;
  margin: 2em 0;
  color: #5d627b;
  background: white;
  border-top: solid 5px #f0ad4e;
  box-shadow: 0 3px 5px rgba(0, 0, 0, 0.22);
}
.box-wa p {
  margin: 0; 
  padding: 0;
}
  
</style>
