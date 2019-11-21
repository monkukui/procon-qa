<template>
  <div>
    <container fluid>
      <v-row>
        <v-col col="12" sm="4">
          <v-tabs
            fixed-tabs
            background-color="transparent"
          >
            <v-tab @click="click(1)"
              class="tab"
            >
              <div
                class="display-100 flex-grow-1 text-center"
              >
                <div style="color: rgb(66, 66, 66);" class="font-weight-bold">
                  新着
                </div>
              </div>
            </v-tab>
            <v-tab @click="click(2)">
              <div
                class="display-100 flex-grow-1 text-center"
              >
                <div style="color: rgb(66, 66, 66);" class="font-weight-bold">
                  いいね
                </div>
              </div>
            </v-tab>
          </v-tabs>
          <v-divider />
        </v-col>
      </v-row>
    </container>
  <br>
  </div>
</template>


<script lang="ts">
import { Vue, Component, Prop, Emit } from 'vue-property-decorator';

@Component
export default class AnswerSetting extends Vue {
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
