<template>
  <div class="modal">
    <transition name="modal">
      <div class="modal-mask">
        <div class="modal-wrapper">
          <div class="modal-container">

            <div class="modal-header">
              <slot name="header">
                <a @click="closeModal">×</a>
              </slot>
            </div>

            <div class="modal-body">
              <slot name="body">
                タイトル [必須] <span style="color: #dc143c;">{{ errorMessageForTitle }}</span>
                <input v-model="questionData.title"></input>
                <span class="num-of-letter">{{ questionData.title.length }} 文字</span>
                <br>
                <br>
                内容 [必須] <span style="color: #dc143c;">{{ errorMessageForBody }}</span>
                <textarea v-model="questionData.body"></textarea>
                <span class="num-of-letter">{{ questionData.body.length }} 文字</span>
                <br>
                <br>
                状態 [任意]
                <select v-model="questionData.state">
                  <option v-for="op in options">{{op}}</option>
                </select>
                URL [任意]
                <input v-model="questionData.url"></input>

              </slot>
            </div>

            <div class="modal-footer">
              <slot name="footer">
                <center>
                  <OriginalButton
                    contents="送信"
                    @click="sendQuestion"
                  />
                </center>
              </slot>
            </div>
          </div>
        </div>
      </div>
    </transition>
  </div>
</template>

<script lang="ts">
import {Component, Emit, Prop, Vue} from 'vue-property-decorator';
import OriginalButton from '@/components/OriginalButton.vue';

interface QuestionData {
  title: string;
  body: string;
  state: string;
  url: string;
}


@Component({
  components: {
    OriginalButton,
  },
})

export default class QuestionModal extends Vue {

  private questionData: QuestionData = {
    title: '',
    body: '',
    state: '',
    url: '',
  };

  private options: string[] = [
    'CE',
    'MLE',
    'TLE',
    'RA',
    'OLE',
    'IE',
    'WA',
    'WJ',
    'WR',
    'AC',
  ];

  private minTitleLength: number = 1;
  private maxTitleLength: number = 30;
  private minBodyLength: number = 1;
  private maxBodyLength: number = 500;

  private errorMessageForTitle: string = '';
  private errorMessageForBody: string = '';

  private validate(): boolean {
    let isValid: boolean = true;
    // タイトルの文字数チェック
    if (this.questionData.title === '') {
      isValid = false;
      this.errorMessageForTitle = 'この項目は必須です';
    } else if (this.questionData.title.length < this.minTitleLength) {
      isValid = false;
      this.errorMessageForTitle = '文字数が足りていません';
    } else if (this.questionData.title.length > this.maxTitleLength) {
      isValid = false;
      this.errorMessageForTitle = '文字数が多すぎます';
    }

    // 内容の文字数チェック
    if (this.questionData.body === '') {
      isValid = false;
      this.errorMessageForBody = 'この項目は必須です';
    } else if (this.questionData.body.length < this.minBodyLength) {
      isValid = false;
      this.errorMessageForBody = '文字数が足りていません';
    } else if (this.questionData.body.length > this.maxBodyLength) {
      isValid = false;
      this.errorMessageForBody = '文字数が多すぎます';
    }
    return isValid;
  }


  // 親コンポーネントに発火
  // @closeModal で呼び出す
  @Emit('closeModal')
  private closeModal(): void {
    // TODO 仕方なくない...? (2 回目)
  }

  private sendQuestion(): void {
    if (this.validate()) {
      // TODO POST の api を叩く
      this.closeModal();
    }
  }
}
</script>

<style>

.num-of-letter {
  font-size: 75%;
}

input {
  font-size:  18px;
  width: 100%;
  height: 100%;
}

select {
  font-size:  18px;
  width: 100%;
  height: 100%;

}

textarea {
  font-size: 18px;
  width: 100%;
  height: 180px;
  resize: none;
}

.modal {
  font-size: 75%;
}

.modal-mask {
  position: fixed;
  z-index: 9998;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, .7);
  display: table;
  transition: opacity .3s ease;
}

.modal-wrapper {
  display: table-cell;
  vertical-align: middle;
}

.modal-container {
  width: 40%;
  height: 96%;
  margin: 0px auto;
  padding: 20px 30px;
  background-color: #fff;
  border-radius: 50px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, .33);
  transition: all .3s ease;
  font-family: Helvetica, Arial, sans-serif;
}

.modal-header h3 {
  margin-top: 0;
  color: #42b983;
}

.modal-body {
  text-align: left;
  margin: 10px 0;
}

.modal-footer {
  margin-top: 40px;
}

.modal-default-button {
  float: right;
}

/*
 * The following styles are auto-applied to elements with
 * transition="modal" when their visibility is toggled
 * by Vue.js.
 *
 * You can easily play with the modal transition by editing
 * these styles.
 */

.modal-enter {
  opacity: 0;
}

.modal-leave-active {
  opacity: 0;
}

.modal-enter .modal-container,
.modal-leave-active .modal-container {
  -webkit-transform: scale(1.1);
  transform: scale(1.1);
}
</style>

