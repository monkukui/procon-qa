<template>
  <div>
    comment
    uid: {{ uid }}
    body: {{ body }}
    date: {{ date }}
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator';

@Component({
  components: {
  },
})
export default class CommentPanel extends Vue {
  @Prop()
  private uid!: string;
  @Prop()
  private body!: string;
  @Prop()
  private date!: string;
}
</script>
