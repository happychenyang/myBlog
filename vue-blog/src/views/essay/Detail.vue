<template>
    <div class="div-detail">
      <h1>Title</h1>

      <div v-html="article" id="content"></div>
    </div>
</template>

<script>

import {defineComponent} from "vue";
import { marked } from "marked"

export default defineComponent({
  name: "detail",
  setup() {
    // marked 选项
    marked.setOptions({
      pedantic: false,
      gfm: true,
      breaks: false,
      sanitize: false,
      smartLists: true,
      smartypants: false,
      xhtml: false,
    });
    const htmlStr = marked('Marked - Markdown Parser\n' +
        '========================\n' +
        '\n' +
        '[Marked] lets you convert [Markdown] into HTML.  Markdown is a simple text format whose goal is to be very easy to read and write, even when not converted to HTML.  This demo page will let you type anything you like and see how it gets converted.  Live.  No more waiting around.\n' +
        '\n' +
        'How To Use The Demo\n' +
        '-------------------\n' +
        '\n' +
        '1. Type in stuff on the left.\n' +
        '2. See the live updates on the right.\n' +
        '\n' +
        'That\'s it.  Pretty simple.  There\'s also a drop-down option above to switch between various views:\n' +
        '\n' +
        '- **Preview:**  A live display of the generated HTML as it would render in a browser.\n' +
        '- **HTML Source:**  The generated HTML before your browser makes it pretty.\n' +
        '- **Lexer Data:**  What [marked] uses internally, in case you like gory stuff like this.\n' +
        '- **Quick Reference:**  A brief run-down of how to format things using markdown.\n' +
        '\n' +
        'Why Markdown?\n' +
        '-------------\n' +
        '\n' +
        'It\'s easy.  It\'s not overly bloated, unlike HTML.  Also, as the creator of [markdown] says,\n' +
        '\n' +
        '> The overriding design goal for Markdown\'s\n' +
        '> formatting syntax is to make it as readable\n' +
        '> as possible. The idea is that a\n' +
        '> Markdown-formatted document should be\n' +
        '> publishable as-is, as plain text, without\n' +
        '> looking like it\'s been marked up with tags\n' +
        '> or formatting instructions.\n' +
        '\n' +
        'Ready to start writing?  Either start changing stuff on the left or\n' +
        '[clear everything](/demo/?text=) with a simple click.\n' +
        '\n' +
        '[Marked]: https://github.com/markedjs/marked/\n' +
        '[Markdown]: http://daringfireball.net/projects/markdown/\n')
    return {
      article: htmlStr
    }
  }
});
</script>

<style scoped>

.div-detail {
  width: 60%;
  margin: 1% auto;
}

</style>