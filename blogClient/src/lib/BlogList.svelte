<script>
    import {deleteMethod, get, post} from '../serverCalls/axiosWrapper.js'
    import { blogPostList } from '../stores/blogPostStore.js'
    import {onMount} from 'svelte'

    onMount(async () =>{
        const result = await get('posts')
        blogPostList.set(result.data)

    })

   $: list = $blogPostList
   const deleteMe = async (id) =>{
    
   }

</script>

<div>
    <h3>List of Blogs</h3>

	
    {#each list as post }
        <div>
            <h4>{post.title}</h4>
            <p>{post.body}</p>
            <button on:click={()=> deleteMe(post.id)}>delete this post </button> 
        </div>
        
    {/each}      
  

</div>

<style>

</style>
