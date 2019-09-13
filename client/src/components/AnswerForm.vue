<template>
  <div class="answer-form">
    <input v-model="text" />
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
  
  private questionId: number;
  private text: string = '';  
  
  private created(): void {
    this.questionId = Number(this.$route.query['questionId']);
  }

  private getToken(): void {
    return localStorage.getItem('token');
  }
  private answer(): void {
    const url = 'api/answers';
    const method = 'POST';
    const headers = {
      'Authorization': `Bearer ${this.getToken()}`,
      'Content-Type': 'application/json; charset=UTF-8'
    }

    const body = JSON.stringify({
      body: this.text,
      favo: 0,
      date: 'yyyy/mm/dd-hh/mm/ss',
      qid: this.questionId,
    });
    
    fetch(url, {method, headers, body}).then(response => {
      if(response.ok) {
        return response.json();
      }
    }).then(json => {
      if(typeof json === 'undefined') {
        return;
      }
    });

  }
}
</script>
