$(document).ready(function(){       
  console.log('asfafawfaw');             
  var canvas = document.getElementById('canvas');
  var stage = new JTopo.Stage(canvas);
  //显示工具栏
  showJTopoToobar(stage);
  var scene = new JTopo.Scene();
  stage.add(scene);
  scene.background = './img/bg.jpg';
  
  var cloudNode = new JTopo.Node('root');
  cloudNode.setSize(30, 26);
  cloudNode.setLocation(360,230);            
  cloudNode.layout = {type: 'circle', radius: 150};
  
  scene.add(cloudNode);
  
  for(var i=1; i<4; i++){
      var node = new JTopo.CircleNode('host' + i);
      node.fillStyle = '200,255,0';
      node.radius = 15;
      node.setLocation(scene.width * Math.random(), scene.height * Math.random());
      node.layout = {type: 'circle', radius: 80};
      
      scene.add(node);                                
      var link = new JTopo.Link(cloudNode, node);
      scene.add(link);
      
      for(var j=0; j<6; j++){
          var vmNode = new JTopo.CircleNode('vm-' + i + '-' + j);
          vmNode.radius = 10;
          vmNode.fillStyle = '255,255,0';
          vmNode.setLocation(scene.width * Math.random(), scene.height * Math.random());
          scene.add(vmNode);                                
          scene.add(new JTopo.Link(node, vmNode));                            
      }
  }
  JTopo.layout.layoutNode(scene, cloudNode, true);
  
  scene.addEventListener('mouseup', function(e){
      if(e.target && e.target.layout){
          JTopo.layout.layoutNode(scene, e.target, true);    
      }                
  });
});