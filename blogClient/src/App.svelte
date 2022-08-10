<script>
import BlogList from "./lib/BlogList.svelte";
import PostEntry from "./lib/PostEntry.svelte";
import PageWrapper from "./lib/PageWrapper.svelte";
import { deleteMethod } from "./serverCalls/axiosWrapper";

const clearPosts= async ()=>{
  await deleteMethod()
}

</script>

<PageWrapper>
  <h1>Local Blog</h1>

  <BlogList />
  <PostEntry />
  <br />
  <button on:click={clearPosts}>Clear All posts</button>
</PageWrapper>
