# Structure

Once you unzip the file, the final name for this project would be `finalProj4Program`. The inner structure is organized as below:



## Directory

### data

This directory would store all the input images, from A - Z (I,O excluded), each symbol contains 100 different figures segmented from the driver license plates. In other words, the dataset contains 2400 (24*100) figures in total. For each single image (jpeg format), the file name follows the same rule that the first character suggests the actual label and the following number indicates the index under this character



### model

This directory would contain the saved version of training model information, either from Python based neural network written in TensorFlow or from the Golang based three layer neural network written from scratch. In specific, `my_model` or `mymodel` subdirectory is generated from TensorFlow, while the `hiddenWeights.model` and `outputWeights.model` are the saved binary file, recording the hidden and output layer of the neural network written from scratch in Golang. This is widely used in the prediction process, which means there is no need to train the model again, but directly load the saved model weights to conduct the prediction on test dataset.



### plot

This directory would contain the confusion matrix generated from Python based image classification. The other part is from the Golang based model. During the direct prediction, I would save the file into this direction where the prediction label is not matched to the actual label, the file is saved in the following format, given actual label is T and the predicted label is G, we could find that `actualLabel_T_predictedLabel_G`.



### python

This directory contains one single py file to utilize the input data and make prediction by TensorFlow. I would like to see whether my neural network written from scratch could achieve same performance as the existing package.





# How to use

main.go is the entrance of the whole project. In this case, I utilized the `flag` method to control the input. You could use `./finalProj4Prgram -help`  in the command line to see the available options



There are three options in this project, that is `epoch`, `file`, `option` and the specific explanation would be suggested as below:



`option:` Three options would be allowed in this project, including `train`, `predict`, `singleImagePred`.

If your option is `train`, the program would automatically read all the images stored in the `data` directory, converted and normalized the image to obtain float matrix, storing the normalized pixel information. Then it would separate the dataset into training and testing dataset, with the ratio 8:2. Option `train` would create a three-layer neural network, including input layer, hidden layer and output layer. Based on the `epoch` specified by user, we could obtain and store the hidden and output weights after training.

The output in the terminal would output each epoch, indicating the Mean Squared Error and the time elapsed. The typical example would suggested as below:

```
You are ready to train the model
The number of epochs used in the training:  5
epoch 1 finished!  MSE:0.3927632258997967  time elapsed: 3.896417658s
epoch 2 finished!  MSE:0.10053819912674045  time elapsed: 3.863140438s
epoch 3 finished!  MSE:0.045825010989174735  time elapsed: 3.855562114s
epoch 4 finished!  MSE:0.02754561135507347  time elapsed: 3.928915654s
epoch 5 finished!  MSE:0.017780493506913013  time elapsed: 4.246954852s
```



If your option is `predict` , the program would also conduct the read, normalize and separate the dataset procedure as `train` option do. But in this case, the model would utilize the testing dat and load the weights stored in previous `train` process. The program would indicate the overall accuracy in the prediction of testing data.

```
======================Overall model analysis======================
The number of test case: 480
The number of perfect match is: 477
accuracy is: 0.99375
```

Moreover, it would indicate the misclassified image in the terminal like this:

```
===================
The actual label is : Y
The predicted label index is : X
The actual image is saved under the plot direction
```

The program would firstly deleted the existing image in `plot` directory and the misclassified images would output to the `plot` directory.

If your option is `singleImagePred`, you have to specify the `file` option in the terminal. It would read the image directory specified by the user in `file` option and load the saved model weights to predict the single image. The default value of `file` is `data/C1.jpeg`

It would output the probability of each 24 label and extract the label with highest probability to be the prediction label. It would also indicate the comparison between input label and prediction label.

```
⎡0.00025700911604224856⎤
⎢  0.005297268909675663⎥
⎢    0.9400345075731126⎥
⎢  0.032581294218555656⎥
⎢ 0.0006765975396444625⎥
⎢ 0.0013419467061111599⎥
⎢  0.057733150699671526⎥
⎢ 0.0005372137239908212⎥
⎢  0.005395146609966048⎥
⎢  0.004006924835547724⎥
⎢  0.006937092469612254⎥
⎢ 0.0018562906196327921⎥
⎢0.00032526741324706224⎥
⎢  0.006799999537798627⎥
⎢  0.014761188585583403⎥
⎢  0.002948870517637401⎥
⎢  0.005647028062065116⎥
⎢  0.001758031997675012⎥
⎢  0.004185116725125173⎥
⎢ 0.0007469257432089711⎥
⎢  0.005817069144051785⎥
⎢ 0.0028969890341807333⎥
⎢ 0.0006008748598513381⎥
⎣    0.0113239804553399⎦
The predicted label is: C
The input image label is: C
```

`epoch:` This would determine the number of epoch used in the training process. The default value would be set as 5 epoch

`file:` This would specify the name of any single PNG file in the data direction based on the saved model weights. The default value is `data/C1.jpeg`