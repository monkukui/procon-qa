<template>
  <div class="pagi-nation">
    <!--hr>
    <p>curPageId: {{ curPageId }}</p>
    <p>totalItemNum: {{ totalItemNum }}</p>
    <p>totalPageNum: {{ totalPageNum }}</p>
    <p>numberOfDisplay: {{ numberOfDisplay }}</p>
    <p>pageIdArray: {{ pageIdArray }}</p>
    <hr-->
    <div class="item">
      <button @click="changePageId(curPageId - 1)"> ( </button>
    </div>
    <div v-for="id in pageIdArray">
      <div v-if="id == curPageId" class="item">
        <button class="active">{{ id }}</button>
      </div>
      <div v-else class="item">
        <button @click="changePageId(id)">{{ id }}</button>
      </div>
    </div>
    <div class="item">
      <button @click="changePageId(curPageId + 1)"> ) </button>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue, Emit } from 'vue-property-decorator';

@Component
export default class PagiNation extends Vue {

  public totalPageNum: number = 0;
  public pageIdArray: number[] = [];

  public numberOfDisplay: number = 7;
  @Prop()
  private curPageId!: number; // 現在のページ番号
  @Prop()
  private totalItemNum!: number; // アイテムの合計数

  @Emit('changePageId')
  public changePageId(id: number): number {
    if (id <= 0) {
      id = 1;
    }
    if (id >= this.totalPageNum) {
      id = this.totalPageNum;
    }
    return id;
  }

  public mounted(): void {
    this.totalPageNum = Math.ceil(this.totalItemNum / this.numberOfDisplay);
    this.pageIdArray = [];
    for (let i: number = 1; i <= this.totalPageNum; i++) {
      this.pageIdArray.push(i);
    }
  }
}
</script>

<style scoped>
button.active {
  background-color: #ccccff;
}
.pagi-nation {
  display: flex;
}
.pagi-nation .item {
}
</style>
