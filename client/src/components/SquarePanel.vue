<template>
  <div>
    <v-item v-slot:default="{ active, toggle }">
      <v-card
        class="d-flex align-center"
        :color="color"
        height="60"
        width="60"
        @click="toggle"
      >
        <v-scroll-y-transition>
          <div
            class="display-100 flex-grow-1 text-center"
          >
            <div style="color: white;"  md="4">
              {{ message }}
              {{ num }}
            </div>
          </div>
        </v-scroll-y-transition>
      </v-card>
    </v-item>
  </div>
</template>


<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator';

@Component({
  componens: {
  },
})
export default class SquarePanel extends Vue {
  @Prop()
  private num: number;
  @Prop()
  private message: string;

  private color: string = 'blue-grey';

  private created(): void {
    if(this.message == '解決済') this.color = '#5cb85c';
    else if(this.message == '未解決') this.color = '#f0ad4e';
  }
}
</script>

<style>

</style>
