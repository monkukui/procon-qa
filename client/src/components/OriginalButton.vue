<template>
  <div class="original-button">
    <a @click="click" href="javascript:void(0);" >{{ contents }}</a>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue, Emit } from 'vue-property-decorator';

@Component({})
export default class PagiNation extends Vue {
  @Prop()
  private contents!: string; // ボタンの上に表示させる文字列

  // 親コンポーネントに発火する
  // 親コンポーネントでは @click="myMethod" のようにする
  @Emit()
  private click(): void {
    // TODO block is empty っていう warn が出るけど仕方なくない...?
  }
}
</script>

<style scoped>
.original-button {
  font-size: 75%;
}

a {
  display: inline-block;
  padding: 0.5em 1em;
  text-decoration: none;
  background: #668ad8;/*ボタン色*/
  color: #FFF;
  border-bottom: solid 4px #627295;
  border-radius: 3px;
}
a:active {
  /*ボタンを押したとき*/
  -webkit-transform: translateY(4px);
  transform: translateY(4px);/*下に動く*/
  box-shadow: 0px 0px 1px rgba(0, 0, 0, 0.2);/*影を小さく*/
  border-bottom: none;
}

</style>
