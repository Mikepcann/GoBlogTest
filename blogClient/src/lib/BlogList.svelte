<script>
    import {deleteById, deleteMethod, get, post} from '../serverCalls/axiosWrapper.js'
    import { blogPostList } from '../stores/blogPostStore.js'
    import {onMount} from 'svelte'

    onMount(async () =>{
        const result = await get('posts')
        blogPostList.set(result.data)

    })

   $: list = $blogPostList
   const deleteMe = async (id) =>{

        await deleteById(id)
   }

</script>

<div>
    <h3>List of Blogs</h3>

	
    {#each list as post }
        <div class="blog">
            <h3>{post.title}</h3>
            <p>{post.body}</p>
            <button on:click={()=> deleteMe(post.id)}>delete this post </button> 
        </div>
    {/each}      
  

</div>

<style>

.blog {
    border: 2px solid #fff;
    padding: 10px;
    margin: 15px;
    width: 500px;
    background-color: antiquewhite;
    color: #000;
}
.blog > h3 {
    text-align: center;
}
</style>
