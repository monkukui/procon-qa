<template>
  <div class="tag-search">
    <h1>質問検索</h1>
    <br>
    <v-alert
      border="top"
      colored-border
      type="info"
      elevation="2"
    >
      質問タイトルや質問者のハンドルネームを入力してください．
      文字列の類似度を計算し，上位 10 件を表示します．
    </v-alert>
    <br>
    <v-textarea
      rows="1"
      light
      v-model="text"
      outlined
      no-resize
      label="検索したい文字列を入力... "
    />
    <v-btn
      color="primary"
      class="mr-4"
      @click="getQuestionsWithEditDistance"
    >
      検索
    </v-btn>

    <div v-if="isReady">
      <div v-for="(value, index) in questions" :key=index>
        <!-- answerd とか, questionedTime とかの命名規則を揃える -->
        <!-- 子コンポーネントには QuestionId だけを渡す-->
        <QuestionPanel
          :questionId="value.id"
        />
      </div>
    </div>
    <div v-else class="text-center">
      <v-progress-circular
        :size="100"
        color="primary"
        indeterminate
      ></v-progress-circular>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import QuestionPanel from '@/components/QuestionPanel.vue';
@Component({
  components: {
    QuestionPanel,
  },
})
export default class TagSearch extends Vue {
  private isReady: boolean = true;
  private text: string = '';
  private questions = [];


  // 質問を取得
  private getQuestionsWithEditDistance(): void {
    this.isReady = false;
    const url = '/api/no-auth/questions/editdistance';
    const headers = {
      'Authorization': `Bearer ${this.getToken()}`,
      'Content-Type': 'application/json; charset=UTF-8',
    };
    const method = 'POST';
    const body = JSON.stringify({
      title: this.text,
    });

    fetch(url, {body, headers, method}).then((response) => {
      if (response.ok) {
        return response.json();
      }
      return [];
    }).then((json) => {
      this.questions = [];
      this.questions = json;
      this.isReady = true;
    });
  }
  private getToken(): any {
    return localStorage.getItem('token');
  }
}
</script>

<style scoped>
.tag-search {
  width: 70%;
  margin-left: 5%;
}
.search-window {
}
</style>
