#pragma once

#include "cef_base.h"
#include "include/capi/cef_app_capi.h"
#include "include/capi/cef_render_process_handler_capi.h"
#include "cef_v8handler.h"
#include "utils.h"


  ///
  // Called after a browser has been created. When browsing cross-origin a new
  // browser will be created before the old browser with the same identifier is
  // destroyed. |extra_info| is an optional read-only value originating from
  // cef_browser_host_t::cef_browser_host_create_browser(),
  // cef_browser_host_t::cef_browser_host_create_browser_sync(),
  // cef_life_span_handler_t::on_before_popup() or
  // cef_browser_view_t::cef_browser_view_create().
  ///
  void CEF_CALLBACK on_browser_created(
      struct _cef_render_process_handler_t* self,
      struct _cef_browser_t* browser,
      struct _cef_dictionary_value_t* extra_info){
    DEBUG_CALLBACK("on_browser_created----------------\n");
    // cef_frame_t *frame=browser->get_main_frame(browser);
    // cef_string_t script=getCefString("alert(window.test);alert(ext.myval);alert(window.testfunc());");
    // frame->execute_java_script(frame,&script,frame->get_url(frame),0);
  }


  ///
  // Called before a browser is destroyed.
  ///
  void CEF_CALLBACK on_browser_destroyed(
      struct _cef_render_process_handler_t* self,
      struct _cef_browser_t* browser){
        DEBUG_CALLBACK("on_browser_destroyed\n");
      }


  ///
  // Called immediately after the V8 context for a frame has been created. To
  // retrieve the JavaScript 'window' object use the
  // cef_v8context_t::get_global() function. V8 handles can only be accessed
  // from the thread on which they are created. A task runner for posting tasks
  // on the associated thread can be retrieved via the
  // cef_v8context_t::get_task_runner() function.
  ///
  void CEF_CALLBACK on_context_created(
      struct _cef_render_process_handler_t* self,
      struct _cef_browser_t* browser,
      struct _cef_frame_t* frame,
      struct _cef_v8context_t* context){
        //DEBUG_CALLBACK("on_context_created\n");
        if(cef_v8context_in_context() == 1) {
          cef_v8context_t* cntx = cef_v8context_get_current_context();
          int isValid = cntx->is_valid(cntx);
          if(isValid == 1) {
            cntx->enter(cntx);
            cef_v8value_t *invocation=context->get_global(context);
            const cef_string_t key=getCefString("test");
            const cef_string_t value=getCefString("hello world js");
            cef_v8value_t *helloValue=cef_v8value_create_string(&value);
            invocation->set_value_bykey(invocation,&key,helloValue,V8_PROPERTY_ATTRIBUTE_NONE);

            const cef_string_t funcKey=getCefString("testfunc");
            const cef_string_t funcName=getCefString("testfunc");
            invocation_handler *v8handler=initialize_v8handler();
            cef_v8value_t  *myfunc=cef_v8value_create_function(&funcName,&v8handler->handler);
            invocation->set_value_bykey(invocation,&funcKey,myfunc, V8_PROPERTY_ATTRIBUTE_NONE);
            cntx->exit(cntx);
        }
      }
    }

  ///
  // Called immediately before the V8 context for a frame is released. No
  // references to the context should be kept after this function is called.
  ///
  void CEF_CALLBACK on_context_released(
      struct _cef_render_process_handler_t* self,
      struct _cef_browser_t* browser,
      struct _cef_frame_t* frame,
      struct _cef_v8context_t* context){
        DEBUG_CALLBACK("on_context_released\n");
      }

  ///
  // Called for global uncaught exceptions in a frame. Execution of this
  // callback is disabled by default. To enable set
  // CefSettings.uncaught_exception_stack_size > 0.
  ///
  void CEF_CALLBACK on_uncaught_exception(
      struct _cef_render_process_handler_t* self,
      struct _cef_browser_t* browser,
      struct _cef_frame_t* frame,
      struct _cef_v8context_t* context,
      struct _cef_v8exception_t* exception,
      struct _cef_v8stack_trace_t* stackTrace){
        DEBUG_CALLBACK("on_uncaught_exception\n");
      }

  ///
  // Called when a new node in the the browser gets focus. The |node| value may
  // be NULL if no specific node has gained focus. The node object passed to
  // this function represents a snapshot of the DOM at the time this function is
  // executed. DOM objects are only valid for the scope of this function. Do not
  // keep references to or attempt to access any DOM objects outside the scope
  // of this function.
  ///
  void CEF_CALLBACK  on_focused_node_changed(
      struct _cef_render_process_handler_t* self,
      struct _cef_browser_t* browser,
      struct _cef_frame_t* frame,
      struct _cef_domnode_t* node){
        DEBUG_CALLBACK("on_focused_node_changed\n");
      }

  ///
  // Called when a new message is received from a different process. Return true
  // (1) if the message was handled or false (0) otherwise. It is safe to keep a
  // reference to |message| outside of this callback.
  ///
  int CEF_CALLBACK on_process_message_received_for_render(
      struct _cef_render_process_handler_t* self,
      struct _cef_browser_t* browser,
      struct _cef_frame_t* frame,
      cef_process_id_t source_process,
      struct _cef_process_message_t* message){
        DEBUG_CALLBACK("on_process_message_received_for_render\n");
        //cef_string_userfree_t name=message->get_name(message);
        //const cef_string_t name = getCefString("https://baidu.com");
        //frame->load_url(frame,&name);
        //cef_string_userfree_free(name);
        return 0;
      }

  ///
  /// Called after WebKit has been initialized.
  ///
  void CEF_CALLBACK on_web_kit_initialized(
      struct _cef_render_process_handler_t* self){
        DEBUG_CALLBACK("on_web_kit_initialized\n");
        cef_string_t extension_name=getCefString("v8/ext");
        cef_string_t jsCode=getCefString("var ext;\n \
            if (!ext)\n \
              ext = {};\n \
            (function() {\n \
              ext.myval = 'My Value!';\n \
            })();");
        
        cef_register_extension(&extension_name,&jsCode,NULL);
  }

render_process_handler * initialize_cef_render_process_handler(){
    DEBUG_CALLBACK("initialize_cef_render_process_handler\n");
    render_process_handler *r=calloc(1,sizeof(render_process_handler));
    cef_render_process_handler_t *handler=(cef_render_process_handler_t *)r;
    handler->on_browser_created= on_browser_created;
    handler->on_browser_destroyed= on_browser_destroyed;
    handler->on_context_created= on_context_created;
    handler->on_context_released= on_context_released;
    handler->on_uncaught_exception= on_uncaught_exception;
    handler->on_focused_node_changed= on_focused_node_changed;
    handler->on_process_message_received= on_process_message_received_for_render;
    handler->on_web_kit_initialized=on_web_kit_initialized;
    return r;
}


void initialize_cef_render_process_handler_direct(cef_render_process_handler_t *handler){
    DEBUG_CALLBACK("initialize_cef_render_process_handler\n");
    handler->base.size = sizeof(cef_render_process_handler_t);
    initialize_cef_base_ref_counted((cef_base_ref_counted_t*)handler);
    /*
      ignore this callback,because it will cause unexcepted action
    */
    handler->on_browser_created= on_browser_created;
    handler->on_browser_destroyed= on_browser_destroyed;
    handler->on_context_created= on_context_created;
    handler->on_context_released= on_context_released;
    handler->on_uncaught_exception= on_uncaught_exception;
    // handler->on_focused_node_changed= on_focused_node_changed;
    handler->on_process_message_received= on_process_message_received_for_render;
    handler->on_web_kit_initialized=on_web_kit_initialized;
}