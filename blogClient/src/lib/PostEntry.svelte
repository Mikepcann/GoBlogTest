<script>
    import { blogPostList } from '../stores/blogPostStore';
    import {post, get} from '../serverCalls/axiosWrapper'

    let title = ''
    let body = ''

    const submitValues = async ()=>{

        if (title == '' || body == ''){
            alert('Enter a title AND a body')
            return
        }

       // post the data to the server
       post('posts',{title, body}) 

       const result = await get('posts')
       blogPostList.update(n => result.data)

        title= ''
        body = ''
    }
</script>

<h3>Enter your post here</h3>
<section>
    
    <label for="title">Title: </label>
        <input type="text" name="title" bind:value={title}>
   

    <label for="title">Body: </label>
    <textarea  name="body" bind:value={body} cols="50" rows="10" maxlength="200"/>   
   
</section>
<button on:click={submitValues}>Submit the blog</button>

<style>
    section {
        display: flex;
        flex-direction: column;
        place-items: center;
    }
    
</style>
