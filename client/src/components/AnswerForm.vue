<template>
  <div class="answer-form">
    
    <div class="mavon-editor">
      <mavon-editor
        :toolbars="markdownOption"
        v-model="text"
        placeholder='回答'
        :boxShadow="fa"
        :ishljs="fa"
      />
    </div>
    <v-btn
      color="primary"
      @click="answer"
      x-large

    >回答する</v-btn>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';

@Component({
  components: {
  },
})
export default class AnswerForm extends Vue {
  

  private markdownOption: any = {
    bold: true,
    italic: true,
    header: true,
    underline: true,
    strikethrough: true,
    mark: true,
    quote: true,
    ol: true,
    ul: true,
    code: true,
    table: true,
    help: true,
    alignleft: true,
    aligncenter: true,
    alignright: true,
    subfield: true,
    preview: true,
    // false
    link: false,
    imagelink: false,
    superscript: false,
    subscript: false,
    undo: false,
    redo: false,
    fullscreen: false,
    readmodel: false,
    htmlcode: false,
    trash: false,
    save: false,
    navigation: false,
  };

  private questionId!: number;
  private text: string = '';

  private created(): void {
    this.questionId = Number(this.$route.query.questionId);
  }

  private getToken(): string | null {
    return localStorage.getItem('token');
  }
  private answer(): void {
    const url = 'api/answers';
    const method = 'POST';
    const headers = {
      'Authorization': `Bearer ${this.getToken()}`,
      'Content-Type': 'application/json; charset=UTF-8',
    };

    const body = JSON.stringify({
      body: this.text,
      favo: 0,
      date: 'yyyy/mm/dd-hh/mm/ss',
      qid: this.questionId,
    });

    fetch(url, {method, headers, body}).then((response) => {
      if (response.ok) {
        return response.json();
      }
    }).then((json) => {
      if (typeof json === 'undefined') {
        return;
      }
    });

    location.reload();

  }
}
</script>

<style scoped>
.mavon-editor {
  z-index: 0;
}
</style>
