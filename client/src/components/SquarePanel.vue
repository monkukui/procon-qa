<template>
  <div>
    <v-item v-slot:default="{ active, toggle }">
      <v-list-item-avatar
        class="d-flex align-center"
        size="60"
        :color="color"
      >
        <v-scroll-y-transition>
          <div
            class="display-100 flex-grow-1 text-center"
          >
            <div style="color: white;" class="font-weight-bold">
              {{ message }}
              {{ num }}
            </div>
          </div>
        </v-scroll-y-transition>
      </v-list-item-avatar>
    </v-item>
  </div>
</template>


<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator';

@Component
export default class SquarePanel extends Vue {
  @Prop()
  private num!: number;
  @Prop()
  private message!: string;
  @Prop()
  private color!: string;
}
</script>

<style>

</style>
