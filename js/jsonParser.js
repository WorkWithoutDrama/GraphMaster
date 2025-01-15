

const fileUploader = document.getElementById('file-uploader');
//import exampleJsonFile from "jsonformatter.json"


const model = `{
  "A1":{
   "value": 7,
   "links": "http//",
   "tags":["@1", "@ПЗ"],
   "initial": [],
   "final": ["s1", "s2"],
   "abbreviation": "ПЗ"
  },
  "A2":{
   "value": 5,
   "links": "http//",
   "tags":["@1", "@ПП"],
   "initial": [],
   "final": ["s3", "s4"],
   "abbreviation": "ПП"
  },
  "A3":{
   "value": 4,
   "links": "http//",
   "tags":["@Смотр"],
   "initial": ["s2", "s3"],
   "final": ["s6"],
   "abbreviation": "Смотр"
  },
  "A4":{
   "value": 8,
   "links": "http//",
   "tags":["@Куп"],
   "initial": [],
   "final": ["s7", "s6"],
   "abbreviation": "Куп"
  },
  "A5":{
   "value": 7,
   "links": "http//",
   "tags":["@Оформ"],
   "initial": [],
   "final": ["s5"],
   "abbreviation": "Оформ"
  },
  "A6":{
   "value": 3,
   "links": "http//",
   "tags":["@Срав"],
   "initial": ["s5", "s6"],
   "final": ["s8"],
   "abbreviation": "Срав"
  }
}`;

//models = JSON.parse(model);
fileUploader.addEventListener('change', (event) => {

  const files = event.target.files;
  console.log(files);
  console.log(typeof files[0]);
  console.log(files[0])
  //const readFile = fs.readFileSync(files[0], 'utf8');
  models = JSON.parse(model);
  //console.log(model);
  //console.log(model[1]);
  fillActions(models);
})

let isa = {};
let fsa = {};
let act = [];

function fillActions(jsonObj) {
  let lstates = [];
  for (var key in jsonObj) {
    //console.log(key);
    let liLast = document.createElement('li');
    liLast.classList.add('brick');
    //liLast.classList.add('action');
    liLast.classList.add('item');
    //liLast.classList.add('click');
    let libutton = document.createElement('button');
    let action = key;
    act[act.length] = action;
    //console.log(jsonObj[key]);
    libutton.innerHTML = action;
    //console.log(jsonObj[0]);
    liLast.append(libutton);
    libutton.classList.add('action__button')
    //libutton.classList.add('click')
    brickslist.append(liLast);

    let opt = document.createElement('option');
      opt.value = action;
      actions_list.append(opt);

    for (var i in jsonObj[key].initial) {
      let is = jsonObj[key].initial[i];
      if (typeof isa[is] === "undefined") {
        isa[is] = [];
        isa[is][0] = action;
        }else{
          isa[is][isa[is].length] = action;
        }


      if (lstates.includes(is)) {}
      else {
        lstates[lstates.length] = is;
      }
  }
    for (var j in jsonObj[key].final) {
      let fs = jsonObj[key].final[j];
      if (typeof fsa[fs] === "undefined") {
        fsa[fs] = [];
        fsa[fs][0] = action;
        }else{
          fsa[fs][fsa[fs].length] = action;
        }
        if (lstates.includes(fs)) {}
        else {
          lstates[lstates.length] = fs;
        }

    }
}


  




for (x in lstates) {
  let opts = document.createElement('option');
  opts.value = lstates[x];
  state_list.append(opts);
}

}

// console.log(isa)
// console.log(fsa)
// console.log(act)

//let brlist = document.getElementById('brickslist');
let brlist = document.getElementsByClassName('click');
console.log(brlist);
for (var key in brlist) {
brlist[key].onclick = function(event) {

  if (event.ctrlKey || event.metaKey) {
    console.log('Спасибо1!');
  } else {
    action = event.target.innerHTML;
    action_name.value = action;
    //console.log(models[action].final);
    action_tags.value = models[action].tags;
    action_link.value = models[action].links;
    action_value.value = models[action].value;

    //console.log(models[action].final[k])
    let flist = document.getElementById('flist');


    listActions(fsa, 'ilist', 'initial', 'initial__state', "i")
    listActions(isa, 'flist', 'final', 'final__state', "f")
  }

}
}

function listActions(stac, nlist, smod, sclass, ind) {

  let list = document.getElementById(nlist);
  while (list.firstChild) {
    list.firstChild.remove();
  }

  for (var k in models[action][smod]) {

    let li = document.createElement('li');
    let inp = document.createElement('input');
    //let h = document.createElement('h4');
    let s = models[action][smod][k];
    inp.placeholder = s;
    inp.setAttribute('list', 'state_list');
    inp.classList.add('item');
    inp.classList.add('state');
    let ids  = ind + k;
    inp.id = ids;
    li.classList.add(sclass);
    //h.append(inp);
    li.append(inp);

    let alu = document.createElement('ul');
    alu.classList.add('ac');
    alu.classList.add('click');

    //let sc = document.createElement('script');
    //sc.innerHTML = "var line = new LeaderLine(document.getElementById('action'),document.getElementById('" + ids + "'),{color: 'black', size: 1.8})";
    //console.log(sc);

   



    for (i in stac[s]) {
      let ac = stac[s][i]
      let ali = document.createElement('li');

      ali.innerHTML = ac;
      ali.classList.add('ac');
      //ali.classList.add('click');
      ali.classList.add('action');
      ali.classList.add('item');
      ali.id = ind + k + "a" + i;
      alu.append(ali)

    }

    li.append(alu);
    //li.append(sc);
    list.append(li);


  }

  let li = document.createElement('li');
  let inp = document.createElement('input');

  inp.setAttribute('list', 'state_list');
  inp.classList.add('item');
  inp.classList.add('state');
  li.classList.add(sclass);

  let alu = document.createElement('ul');
  let ali = document.createElement('action__button');
  li.append(inp);
  //alu.append(ali);
  li.append(alu);

  list.append(li);


}
// console.log(isa)
// console.log(fsa)
// console.log(act)



