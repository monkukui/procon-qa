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

  // FIXME
  private tags: string[] = [
    'AtCoder',
    'ABC',
    'C++',
  ];

  // TODO query じゃなくて, param にしようか...?
  // じゃないと tags: string[] が渡せない
  // いや, tags: string[] を渡す必要はないと思う
  // id だけ渡すことができれば, DB からエイっと検索 !!
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
  
</style>
