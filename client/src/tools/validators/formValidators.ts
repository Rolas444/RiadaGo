 export const validateEmail=(email: string)=> {
        if(email!=''){
          const re = /^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
          //console.log(re.test(String(email).toLowerCase()))
          var emailTest = re.test(String(email).toLowerCase())
          return emailTest;
        }
    
    }


