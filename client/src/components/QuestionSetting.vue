<template>
  <div>
    <v-row align="center">
      <v-col cols="12" sm="6">
        <p class="font-weight-bold display-1">並び替え</p>
      </v-col>
      <v-col class="text-center" cols="12" sm="1">
        <div class="my-2">
          <v-btn text outlined color="primary" @click="click(1)" v-if="curMode===1"><div class="font-weight-bold">新着</div></v-btn>
          <v-btn text outlined @click="click(1)" v-else>新着</v-btn>
        </div>
      </v-col>
      <v-col class="text-center" cols="12" sm="1">
        <div class="my-2">
          <v-btn text outlined color="primary" @click="click(2)" v-if="curMode===2"><div class="font-weight-bold">回答数</div></v-btn>
          <v-btn text outlined @click="click(2)" v-else>回答数</v-btn>
        </div>
      </v-col>

      <v-col class="text-center" cols="12" sm="1">
        <div class="my-2">
          <v-btn text outlined color="primary" @click="click(3)" v-if="curMode===3"><div class="font-weight-bold">閲覧数</div></v-btn>
          <v-btn text outlined @click="click(3)" v-else>閲覧数</v-btn>
        </div>
      </v-col>

      <v-col class="text-center" cols="12" sm="1">
        <div class="my-2">
          <v-btn text outlined color="primary" @click="click(4)" v-if="curMode===4"><div class="font-weight-bold">いいね</div></v-btn>
          <v-btn text outlined @click="click(4)" v-else>いいね</v-btn>
        </div>
      </v-col>

      <v-col class="text-center" cols="12" sm="1">
        <div class="my-2">
          <v-btn text outlined color="primary" @click="click(5)" v-if="curMode===5"><div class="font-weight-bold">未解決</div></v-btn>
          <v-btn text outlined @click="click(5)" v-else>未解決</v-btn>
        </div>
      </v-col>
    </v-row>
  <v-divider />
  <br>
  </div>
</template>


<script lang="ts">
import { Vue, Component, Prop, Emit } from 'vue-property-decorator';

@Component
export default class QuestionSetting extends Vue {
  @Prop()
  private curMode!: number;
  @Emit()
  private click(mode: number): number {
    return mode;
  }
}
</script>

<style>
</style>
